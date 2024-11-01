package transcode

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/LeoYoung-code/cast"
)

// transcodeVideo 函数用于读取本地视频文件并转码
func transcodeVideo(inputFile string, outputFile string) error {
	// 构建 ffmpeg 命令
	// cmd := exec.Command("ffmpeg", "-i", inputFile, "-c:v", "libx264", "-preset", "slow", "-crf", "22", "-c:a", "aac", "-b:a", "128k", outputFile)
	// cmd := exec.Command("ffmpeg", "-i", inputFile, "-c:v", "libx264", "-preset", "ultrafast", "-crf", "23", "-c:a", "aac", "-b:a", "128k", outputFile)
	// cmd := exec.Command("ffmpeg", "-i", inputFile, "-c:v", "libx264", "-preset", "veryfast", "-crf", "22", "-c:a", "aac", "-b:a", "128k", outputFile)
	cmd := exec.Command("ffmpeg", "-i", inputFile, "-c:v", "libx264", "-preset", "veryfast", outputFile)

	// 显示转码过程中的输出
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("开始转码视频：%s\n", inputFile)
	start := time.Now()
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("转码失败: %v", err)
	}
	fmt.Printf("转码完成，耗时: %v\n", time.Since(start))
	return nil
}

func do() {
	// 输入文件路径
	inputFile := "/Users/staff/Dream/utils/input.mp4" // 替换为您的输入视频文件路径
	// 输出文件路径
	outputFile := filepath.Join("/Users/staff/Dream/utils/output.mp4") // 设置转码后的输出文件路径

	// 检查输入文件是否存在
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		fmt.Printf("输入文件不存在: %s\n", inputFile)
		return
	}

	// 调用转码函数
	err := transcodeVideo(inputFile, outputFile)
	if err != nil {
		fmt.Println("转码过程出错:", err)
	} else {
		fmt.Printf("视频转码成功！输出文件位于: %s\n", outputFile)
	}
}

// transcodeVideoIO 使用输入 io.Reader 并返回转码完成后的 io.Reader
func transcodeVideoIO(input io.Reader, cmd *exec.Cmd) (io.Reader, error) {
	// 创建一个缓冲区，用于存储 ffmpeg 的输出
	var outputBuffer bytes.Buffer

	// 配置 ffmpeg 的 stdin 和 stdout
	cmd.Stdin = input          // 将输入 io.Reader 设置为标准输入
	cmd.Stdout = &outputBuffer // 将标准输出设置为输出缓冲区
	cmd.Stderr = os.Stderr     // 可选：打印 ffmpeg 的错误信息

	// 执行命令并等待完成
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("转码失败: %v", err)
	}

	// 返回包含输出的 io.Reader
	return &outputBuffer, nil
}

func doIO() {
	// 输入文件路径
	inputFilePath := "/Users/staff/Dream/utils/input.mp4" // 替换为您的输入视频文件路径
	// 输出文件路径
	round := rand.Intn(100) + 1
	outputFilePath := filepath.Join("output" + cast.ToString(round) + ".mp4") // 设置转码后的输出文件路径
	// 打开输入文件（可以是任意 io.Reader）
	inputFile, err := os.Open(inputFilePath) // 替换为源视频文件路径
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer inputFile.Close()

	// 创建并配置 ffmpeg 命令
	// cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-c:v", "libx264", "-preset", "veryfast", "-crf", "28", "-c:a", "aac", "-b:a", "128k", "-f", "mp4", "pipe:1")
	// cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-c:v", "libx264", "-preset", "veryfast", "-f", "mp4", "pipe:1")
	cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-c:v", "libx264", "-preset", "veryfast", "-f", "mp4", outputFilePath)
	// 调用转码函数
	outputReader, err := transcodeVideoIO(inputFile, cmd)
	if err != nil {
		fmt.Println("转码过程出错:", err)
		return
	}

	// 将转码后的数据输出到文件
	outFile, err := os.Create("output.mp4")
	if err != nil {
		fmt.Println("创建输出文件失败:", err)
		return
	}
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(outFile)

	// 将 outputReader 中的数据复制到输出文件
	if _, err := io.Copy(outFile, outputReader); err != nil {
		fmt.Println("写入输出文件失败:", err)
	} else {
		fmt.Println("视频转码成功并保存到 output.mp4！")
	}
}
