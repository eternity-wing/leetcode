package problems

import "fmt"

const Visiting = -1
const Visited = 1
const NotVisited = 0

func canFinish(numCourses int, prerequisites [][]int) bool {
	visitedCourses := make([]int, numCourses)
	prerequisites = getCoursePrerequisitesMap(numCourses, prerequisites)
	for i := 0; i < numCourses; i++ {
		if visitedCourses[i] == NotVisited && !doDfs(i, visitedCourses, prerequisites) {
			return false
		}
	}
	return true

}

func doDfs(course int, visitedCourses []int, prerequisites [][]int) bool {
	if visitedCourses[course] == Visiting {
		return false
	}

	if visitedCourses[course] == Visited {
		return true
	}
	visitedCourses[course] = Visiting

	coursePrerequisites := prerequisites[course]
	for _, prerequisite := range coursePrerequisites {
		if !doDfs(prerequisite, visitedCourses, prerequisites) {
			return false
		}
	}
	visitedCourses[course] = Visited
	return true
}

func getCoursePrerequisitesMap(numCourses int, prerequisites [][]int) [][]int {
	coursePrerequisitesMap := make([][]int, numCourses)
	for _, prerequisitesRecord := range prerequisites {
		course := prerequisitesRecord[0]
		coursePrerequisite := prerequisitesRecord[1]
		coursePrerequisitesMap[course] = append(coursePrerequisitesMap[course], coursePrerequisite)
	}
	return coursePrerequisitesMap
}

func RunCanFinish() {
	prerequisites := [][]int{{1, 0}, {0, 1}}
	fmt.Printf("\n%+v\n", canFinish(2, prerequisites))
}