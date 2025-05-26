package image

import (
	"github.com/TestsLing/aj-captcha-go/const"
	"github.com/TestsLing/aj-captcha-go/util"
	"log"
	"os"
	"path/filepath"
)

var backgroundImageArr []*util.ImageUtil
var clickBackgroundImageArr []*util.ImageUtil
var templateImageArr []*util.ImageUtil

// var resourceAbsPath string // No longer needed globally as font path is passed directly

func SetUp(resourcePath string) {
	// resourceAbsPath = resourcePath // Not needed globally
	root := resourcePath

	backgroundImageRoot := root + constant.DefaultBackgroundImageDirectory
	templateImageRoot := root + constant.DefaultTemplateImageDirectory
	clickBackgroundImageRoot := root + constant.DefaultClickBackgroundImageDirectory

	// Initialize slices to avoid nil pointer issues if directories are empty
	backgroundImageArr = make([]*util.ImageUtil, 0)
	templateImageArr = make([]*util.ImageUtil, 0)
	clickBackgroundImageArr = make([]*util.ImageUtil, 0)

	err := filepath.Walk(backgroundImageRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %s: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		imgUtil := util.NewImageUtil(path)
		if imgUtil == nil {
			log.Printf("Failed to load background image: %s", path)
			return nil // Continue with other files
		}
		backgroundImageArr = append(backgroundImageArr, imgUtil)
		return nil
	})
	if err != nil {
		log.Printf("Error walking background image directory: %v", err)
	}

	err = filepath.Walk(templateImageRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %s: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		imgUtil := util.NewImageUtil(path)
		if imgUtil == nil {
			log.Printf("Failed to load template image: %s", path)
			return nil // Continue with other files
		}
		templateImageArr = append(templateImageArr, imgUtil)
		return nil
	})
	if err != nil {
		log.Printf("Error walking template image directory: %v", err)
	}

	err = filepath.Walk(clickBackgroundImageRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %s: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		imgUtil := util.NewImageUtil(path)
		if imgUtil == nil {
			log.Printf("Failed to load click background image: %s", path)
			return nil // Continue with other files
		}
		clickBackgroundImageArr = append(clickBackgroundImageArr, imgUtil)
		return nil
	})
	if err != nil {
		log.Printf("Error walking click background image directory: %v", err)
	}

	if len(backgroundImageArr) == 0 {
		log.Println("Warning: No background images loaded. Check directory:", backgroundImageRoot)
	}
	if len(templateImageArr) == 0 {
		log.Println("Warning: No template images loaded. Check directory:", templateImageRoot)
	}
	if len(clickBackgroundImageArr) == 0 {
		log.Println("Warning: No click background images loaded. Check directory:", clickBackgroundImageRoot)
	}
}

func GetBackgroundImage() *util.ImageUtil {
	if len(backgroundImageArr) == 0 {
		log.Println("Error: No background images available.")
		return nil
	}
	originalImageUtil := backgroundImageArr[util.RandomInt(0, len(backgroundImageArr)-1)]
	return originalImageUtil.Copy()
}

func GetTemplateImage() *util.ImageUtil {
	if len(templateImageArr) == 0 {
		log.Println("Error: No template images available.")
		return nil
	}
	originalImageUtil := templateImageArr[util.RandomInt(0, len(templateImageArr)-1)]
	return originalImageUtil.Copy()
}

func GetClickBackgroundImage() *util.ImageUtil {
	// Note: Original code used templateImageArr for max length, assuming it's a typo
	// and should use clickBackgroundImageArr
	if len(clickBackgroundImageArr) == 0 {
		log.Println("Error: No click background images available.")
		return nil
	}
	originalImageUtil := clickBackgroundImageArr[util.RandomInt(0, len(clickBackgroundImageArr)-1)]
	return originalImageUtil.Copy()
}
