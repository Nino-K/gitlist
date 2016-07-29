package storage

import (
	"encoding/gob"
	"os"
	"testing"

	"github.com/nino-k/gitlist/apihandler"
)

var (
	storage  *Storage
	testPath = os.TempDir() + "testFile"
)

func TestMain(m *testing.M) {
	//defer cleanUp(testPath)
	storage = New(testPath)
	m.Run()
	//	sig := m.Run()
	//	if sig != 0 {
	//		os.Exit(sig)
	//	}
	cleanUp(testPath)
}

func TestEncode(t *testing.T) {
	t.Log("Always creates a new file and override existing one")
	{
		storage := New(testPath)

		// save some data
		firstRepos := apihandler.GetTestRepos("first_test", "second_test", "third_test")
		dataToBeRemoved := ConvertToStorageData(firstRepos.Items)
		err := storage.Encode(dataToBeRemoved)
		if err != nil {
			t.Errorf("Encode failed: %v", err)
		}

		// save more data
		secondRepos := apihandler.GetTestRepos("one", "two")

		data := ConvertToStorageData(secondRepos.Items)
		err = storage.Encode(data)
		if err != nil {
			t.Errorf("Encode failed: %v", err)
		}

		// make sure only second set of data is retrieved
		var decodedData []Data
		dataFile, err := os.Open(testPath)
		if err != nil {
			t.Errorf("Not Expected: %v", err)
		}
		decoder := gob.NewDecoder(dataFile)
		err = decoder.Decode(&decodedData)
		if err != nil {
			t.Errorf("gob decoder failed: %v", err)
		}

		testRepos := ConvertToRepository(decodedData)

		// validate agains second set
		apihandler.VerifyItems(secondRepos.Items, testRepos, t)
	}
}

func TestDecodeError(t *testing.T) {
	t.Log("Returns an error if")
	{
		t.Log("file does not exist")
		{
			err := os.Remove(testPath)
			if err != nil {
				t.Errorf("File remove failed: %v", err)
			}
			verifyDecodingFailure(storage, t)
		}

		t.Log("underlying gob decoder fails")
		{
			dataFile, err := os.Create(testPath)
			if err != nil {
				t.Errorf("File open error: %v", err)
			}
			encoder := gob.NewEncoder(dataFile)
			err = encoder.Encode("Garbage")
			if err != nil {
				t.Errorf("Encoder failed: %v", err)
			}
			verifyDecodingFailure(storage, t)
		}
	}
}

func TestDecode(t *testing.T) {
	t.Log("Retrieves correct data")
	{
		repos := apihandler.GetTestRepos("one", "two", "three", "four")
		dataToEncode := ConvertToStorageData(repos.Items)
		err := storage.Encode(dataToEncode)
		if err != nil {
			t.Errorf("Encode failed: %v", err)
		}
		verifyDecodingSuccess(storage, repos.Items, t)
	}
}

func TestDecodeMultipleTimes(t *testing.T) {
	t.Log("Can be called multiple times")
	{
		repos := apihandler.GetTestRepos("one", "two", "three", "four")
		dataToEncode := ConvertToStorageData(repos.Items)
		err := storage.Encode(dataToEncode)
		if err != nil {
			t.Errorf("Encode failed: %v", err)
		}

		// decode and verify repos first time
		verifyDecodingSuccess(storage, repos.Items, t)

		// decode and verify repos second time
		verifyDecodingSuccess(storage, repos.Items, t)

		// decode and verify repos third time
		verifyDecodingSuccess(storage, repos.Items, t)
	}
}

func TestGetDataById(t *testing.T) {
	t.Log("Returns single Data that is retrieved by a given Id")
	{
		firstRepoId := 1

		repos := apihandler.GetTestRepos("one")
		dataToEncode := ConvertToStorageData(repos.Items)
		err := storage.Encode(dataToEncode)
		if err != nil {
			t.Errorf("Encode failed: %v", err)
		}
		data, err := storage.GetDataById(0)
		if err != nil {
			t.Errorf("GetDataById failed: %v", err)
		}
		if data.Id != firstRepoId {
			t.Error("Id does not match")
		}

		apihandler.Verify(data.Repo, dataToEncode[0].Repo, t)
	}
}

func TestGetDataByIdError(t *testing.T) {
	t.Log("Returns error and nil data if data not found")
	{
		data, err := storage.GetDataById(1)
		if data != nil {
			t.Error("Expected data to be nil")
		}
		if err == nil {
			t.Error("Expected error not to be nil")
		}
	}
}

func cleanUp(path string) {
	os.Remove(path)
}

func verifyDecodingSuccess(storage *Storage, repos []apihandler.Repository, t *testing.T) {
	decodedRepos, err := storage.Decode()
	if err != nil {
		t.Errorf("Decode failed: %v", err)
	}
	if len(decodedRepos) == 0 {
		t.Error("decodedRepos was not retrieved")
	}

	testRepos := ConvertToRepository(decodedRepos)
	apihandler.VerifyItems(repos, testRepos, t)

}

func verifyDecodingFailure(storage *Storage, t *testing.T) {
	repos, err := storage.Decode()
	if err == nil {
		t.Error("Expected an error from Decode")
	}
	if repos != nil {
		t.Error("expected repos to be nil")
	}
}
