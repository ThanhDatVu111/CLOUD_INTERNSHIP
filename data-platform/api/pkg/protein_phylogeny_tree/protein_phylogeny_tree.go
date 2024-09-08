package pkg

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

// Alignment runs MAFFT on the provided sequences file and writes the output to the alignment file.
// It takes the sequences file name and the alignment file name as parameters and returns an error if any occurs.
func Alignment(sequencesFileName, alignmentFileName string) error {
	// Path to the WSL executable and MAFFT command
	mafftPath := "mafft"

	// Arguments for the MAFFT command
	args := []string{
		"--auto",
		sequencesFileName,
	}


	// Create the WSL command
	cmd := exec.Command(mafftPath, args...)

	// Run the command and capture its output
	output, err := cmd.Output()

	
	if err != nil {
		return fmt.Errorf("failed to execute MAFFT: %w", err)
	}


	// Write the output to the alignment file
	err = ioutil.WriteFile(alignmentFileName, output, 0644)
	if err != nil {
		return fmt.Errorf("failed to write output to file: %w", err)
	}

	return nil
}

// ------------------------------------------------

/*
 input_file = "aligned_raw.fasta"
    output_file = "alignedFormatted.png" # You need to create the output_file yourself
    
    with open(input_file, "r") as infile, open(output_file, "w") as outfile:
        for line in infile:
            if line.startswith(">"):
                # Remove the content after the colon and replace the colon with a space
                line = line.replace(" ", "-").replace(",", "").replace(":", "_").replace("(", "-").replace(")", "-") \
                .replace("[", "-").replace("]", "-")
            outfile.write(line)
    
    print("finish")
*/


// ReadFastaInput reads the input file

// lib.ReadFastInput
// lib.ReadText
// go run Fasta.go
// bash Fasta.sh

/* 
	cmd := exec.Command("yourcommand", "some", "args")
if err := cmd.Run(); err != nil { 
    fmt.Println("Error: ", err)
}   

os.systen('ping google.com')

*/



// func ReadFastaInput(fileName string) (byte, error) {
// 	data := make([]byte, 100)
// 	fileBytes, err := os.open(fileName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return fileBytes, nil
// }

// ProcessFastaInput cleans up the file
// func ProcessFastaInput(){
// 	for _, line := ReadFastaInput(){
// 		// process and clean the line
// 	}

// 	// write it to the cleaned up file
	
// }
