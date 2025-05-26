package test

import (
	"fmt"
	"github.com/TestsLing/aj-captcha-go/service"
	"github.com/TestsLing/aj-captcha-go/util"
	"image/color"
	"os" // Import os
	"testing"
)

func TestBlockPuzzleCaptchaService_Get(t *testing.T) {
	//
	//vo := &vo2.CaptchaVO{}
	//b := &service.BlockPuzzleCaptchaService{}
	//res := b.Get(*vo)
	//
	//fmt.Println(res)
}

func TestImage(t *testing.T) {
	// Define paths. These should ideally be relative or configurable for robust testing.
	// For now, using absolute paths as in the original test.
	// Ensure these paths are valid in the test environment.
	// It's better to use a test resource directory.
	// resourcePath := "../resources" // Example: if resources is one level up from test dir
	// imagePath := resourcePath + "/defaultImages/jigsaw/original/1.png"
	// fontPath := resourcePath + "/fonts/WenQuanZhengHei.ttf"

	// Hardcoded paths from the original test, ensure these exist or replace with placeholder/mock.
	// This test will likely fail if these exact paths don't exist in the execution environment.
	imagePath := "/mnt/f/workspace/aj-captcha-go/resources/defaultImages/jigsaw/original/1.png"
	fontPath := "/mnt/f/workspace/aj-captcha-go/resources/fonts/WenQuanZhengHei.ttf"

	// Check if imagePath exists, skip if not
	if _, errStat := os.Stat(imagePath); os.IsNotExist(errStat) {
		t.Skipf("Skipping TestImage: resource file not found at %s", imagePath)
	}
	// Check if fontPath exists, skip if not
	if _, errStat := os.Stat(fontPath); os.IsNotExist(errStat) {
		t.Skipf("Skipping TestImage: resource file not found at %s", fontPath)
	}

	backgroundImage := util.NewImageUtil(imagePath)
	if backgroundImage == nil {
		// This case should ideally be covered by the OsStat check, but as a safeguard:
		t.Fatalf("NewImageUtil returned nil, though resource file was expected at: %s", imagePath)
	}

	// 为背景图片设置水印
	err := backgroundImage.SetText(fontPath, "牛逼AAA", 14, color.RGBA{R: 120, G: 120, B: 255, A: 255})
	if err != nil {
		t.Fatalf("SetText failed: %v. Ensure font path is correct: %s", err, fontPath)
	}
	backgroundImage.DecodeImageToFile() // This saves an image, useful for debugging.
}

func TestIntCovert(t *testing.T) {

	cache := service.NewMemCacheService(10)

	cache.Set("test1", "tes111", 0)

	val := cache.Get("test1")

	fmt.Println(val)

}
