package handler

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Section struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Lessons []Lesson `json:"lessons"`
}

type Lesson struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
}

func LessonListHandler(c *fiber.Ctx) error {

	var lessons []Section

	root := "./lessons"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// var curSection string

		switch strings.Count(path, "\\") {
		case 0:
			return nil
		case 1:
			// sections
			parts := strings.Split(path, "\\")
			name := parts[len(parts)-1]
			nparts := strings.Split(name, "-")
			id, err := strconv.Atoi(nparts[0])
			if err != nil {
				pan := fmt.Sprintf("THERE ARE RULES! FOLLOW THEM! Can't convert %s to type int.", nparts[0])
				panic(pan)
			}
			lessons = append(lessons, Section{
				ID:      id,
				Name:    nparts[1],
				Lessons: []Lesson{},
			})
			// curSection = nparts[1]
			return nil
		case 2:
			// lessons
			parts := strings.Split(path, "\\")
			name := parts[len(parts)-1]
			nparts := strings.Split(name, "-")
			id, err := strconv.Atoi(nparts[0])
			if err != nil {
				pan := fmt.Sprintf("THERE ARE RULES! FOLLOW THEM! Can't convert %s to type int.", nparts[0])
				panic(pan)
			}

			lesson := Lesson{
				ID:   id,
				Name: nparts[1],
				Text: "",
			}
			pos := len(lessons) - 1

			lessons[pos].Lessons = append(lessons[pos].Lessons, lesson)

			return nil
		case 3:
			// files
			parts := strings.Split(path, "\\")
			name := parts[len(parts)-1]

			if name == "text.md" {
				secpos := len(lessons) - 1
				lespos := len(lessons[secpos].Lessons) - 1
				lessons[secpos].Lessons[lespos].Text = path
			}
			return nil
		default:
			return errors.New("dini Ordnerstruktur isch unbruchbar, fix mau")
		}
	})
	if err != nil {
		panic(err)
	}

	return c.JSON(lessons)
}
