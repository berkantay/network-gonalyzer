package input

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func SelectInterface(interfaceLength int) (int, error) {

	fmt.Println("Please select interface:")

	var interface_number int

	fmt.Scanf("%d", &interface_number)

	if interface_number+1 > interfaceLength {

		return interface_number, errors.New("interface number error")
	}

	fmt.Println("Selected interface option is:", interface_number)

	return interface_number, nil
}

func openLocalIpFile(filePath string) *os.File {

	content, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	return content

}
