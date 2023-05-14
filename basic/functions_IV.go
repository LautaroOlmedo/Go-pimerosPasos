package basic

import "fmt"

// RECURSIVIDAD

func FunctionsIV(runner uint8, totalRunners uint8){
	fmt.Println("The runner", runner, "starts the race")
	otherRunner := totalRunners > runner

	if(otherRunner){
		nextRunner := runner + 1
		fmt.Println("Runner", runner, "hands the ball to runner", nextRunner)
		FunctionsIV(nextRunner, totalRunners)
	}

	fmt.Println("The runner", runner, "ends the race")
}

func FunctionsIV2(MountainEchoMsg string, itrator uint8){
	if(itrator > 1){
		FunctionsIV2(MountainEchoMsg, itrator - 1)
	}

	fmt.Println(MountainEchoMsg, itrator)
}