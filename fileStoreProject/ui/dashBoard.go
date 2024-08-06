package ui

import (
	"file/component"
	"file/models"
	"file/util"
	"fmt"
)

func ShowAssignCourse() {
	fmt.Println("--------------------------Courses-------------------------------")
	var assignCourse models.Course
	util.LoadCourses(&assignCourse)
	for _, c := range assignCourse.Courses {
		fmt.Println("ID ", c.ID)
		fmt.Println("Title ", c.Title)
		for _, v := range c.Lessons {
			fmt.Print("Lesson ID: ", v.ID)
			fmt.Println(" Lesson Title:", v.Title)
		}
		fmt.Println("----------------------------------------")
	}
}

func getTodo(username string) {

	for {
		fmt.Println("---------------------To Do List--------------------------------------")
		fmt.Println("To view ToDo press 1")
		fmt.Println("To Add in  ToDo press 2")
		fmt.Println("To delete in  ToDo press 3")
		fmt.Println("To go back press 4 ")
		var choice int
		fmt.Println("Enter your choice: ")
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			component.ShowToDo(username)
			fmt.Println("-------------------------------------------")
			break
		case 2:
			component.AddToDo(username)
			fmt.Println("-------------------------------------------")
			break
		case 3:
			component.DeleteToDo(username)
			fmt.Println("-------------------------------------------")
			break
		case 4:
			return
		default:
			fmt.Println("Enter Invalid Choice ")

		}
	}

}
func getProgress(username string) {
	for {
		fmt.Println("---------------------------Progress---------------------------------")
		fmt.Println("To view Progress press 1")
		fmt.Println("To Update in  Progress press 2")
		fmt.Println("To go back press 3 ")
		var choice int
		fmt.Println("Enter your choice: ")
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			component.ShowUserProgress(username)
			fmt.Println("----------------------------------------")
			break
		case 2:
			err := component.UpdateUserProgress(username)
			if err != nil {
				return
			}
			break
		case 3:
			return

		default:
			fmt.Println("Invalid Choice")

		}
	}
}
func ViewProfile(username string) {
	fmt.Println("--------------Profile-----------------------")
	users, err := component.ReadUsers()
	if err != nil {
		fmt.Println("Error reading users:", err)
		return
	}

	user, ok := users[username]
	if !ok {
		fmt.Println("User not found:", username)
		return
	}

	fmt.Printf("Profile for %s:\n", username)
	fmt.Printf("Username: %s\n", user.Username)
	fmt.Printf("Age: %d\n", user.Age)
	fmt.Printf("Mobile: %s\n", user.Mobile)
}
func getDailyStatus(username string) {

	for {
		fmt.Println("--------------Daily Status---------------")
		fmt.Println("For Display press 1")
		fmt.Println("For Update press 2")
		fmt.Println("To go back press 3 ")

		var choice int
		fmt.Println("Enter your choice: ")
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			component.ShowDailyStatus(username)
			fmt.Println("----------------------------------------")
			break
		case 2:
			component.UpdateDailyStatus(username)
			fmt.Println("-----------------------------------------")
			break
		case 3:
			return
		default:
			fmt.Println("Invalid Choice")
		}
	}

}
func DashBoard(username string) {
	//var wg sync.WaitGroup
	fmt.Println("---------------------Dash Board--------------------------------------")

	for {
		fmt.Println("For view assigned courses press 1")
		fmt.Println("For view profile press 2")
		fmt.Println("For Open ToDo press 3")
		fmt.Println("For checking progress press 4")
		fmt.Println("For Daily status press 5")
		fmt.Println("For exit press 6")
		var choice int
		fmt.Println("Enter your choice: ")
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			//wg.Add(1)
			//defer wg.Done()
			ShowAssignCourse()

			fmt.Println("---------------------------------------------")
		case 2:
			//wg.Add(1)
			//defer wg.Done()

			ViewProfile(username)

			fmt.Println("---------------------------------------------")

		case 3:
			//wg.Add(1)
			//defer wg.Done()

			getTodo(username)

			fmt.Println("---------------------------------------------")

		case 4:
			//wg.Add(1)
			//defer wg.Done()

			getProgress(username)

			fmt.Println("---------------------------------------------")
			break
		case 5:
			//wg.Add(1)
			//defer wg.Done()

			getDailyStatus(username)

			fmt.Println("---------------------------------------------")
		case 6:
			return
		default:
			fmt.Println("Enter Invalid Choice ")

		}
	}

}
