package pkg

import (
	"testing"
	"io/ioutil"
	"os"
)


// TestAlignment is a test function for the Alignment function.
func TestAlignment(t *testing.T) {
	// Define input and output file names
	sequencesFileName := "testData/sequences.fasta"
	alignmentFileName := "testData/aligned_raw.fasta"

	// Ensure the sequences.fasta file exists for the test
	if _, err := os.Stat(sequencesFileName); os.IsNotExist(err) {
		t.Fatalf("input file %s does not exist", sequencesFileName)
	}

	// Run the Alignment function
	err := Alignment(sequencesFileName, alignmentFileName)
	if err != nil {
		t.Fatalf("Alignment failed: %v", err)
	}

	// Check if the output file was created
	if _, err := os.Stat(alignmentFileName); os.IsNotExist(err) {
		t.Fatalf("output file %s was not created", alignmentFileName)
	}

	// Clean up: remove the output file after the test
	defer os.Remove(alignmentFileName)

	// Optionally, read and check the contents of the output file
	content, err := ioutil.ReadFile(alignmentFileName)
	if err != nil {
		t.Fatalf("failed to read output file %s: %v", alignmentFileName, err)
	}

	if len(content) == 0 {
		t.Fatalf("output file %s is empty", alignmentFileName)
	}

	// Print success message if everything passes
	t.Log("Success")
}











// ------------------------------------------------------

// ReadFastaInputTest test
// func ReadFastaInputTest(t *tesing.Testing) error {
//  	smallInputPutFile := 'testData/<file_name>'
// 	fileBytes, err := ReadFastaInput(inputPutFile)

// 	if err != nil{
// 		t.error("error reading %s", smallInputPutFile)
// 	}

// 	assert.NotEqual(t, len(fileBytes), 0, "File has data")

// }

// // ReadFastaInputTest test
// func ReadFastaInputTest(t *tesing.Testing) error {
//    largeInputPutFile := 'testData/<file_name>'
//    fileBytes, err := ReadFastaInput(inputPutFile)

//    if err != nil{
// 	   t.error("error reading %s", largeInputPutFile)
//    }

//    assert.NotEqual(t, len(fileBytes), 0, "File has data")

// }




