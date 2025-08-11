// В чем проблема? Утечка памяти
// срез строки v[:100], Go не создает новую строку с новыми данными в памяти
// Вместо этого он создает новый заголовок строки (string header), который указывает на тот же
// самый базовый массив байт, что и исходная большая строка

// 1. someFunc создается огромная строка v
// 2. justString получает срез этой строки
// 3. someFunc завершается, и локальная переменная v выходит из области видимости
//  маленькая 100-байтовая строка не дает сборщику мусора освободить целый килобайт

package main

import (
	"strings"
)

var justString string

func createHugeString(size int) string {
	return strings.Repeat("a", size)
}

func someFunc() {
	v := createHugeString(1 << 10) // 1024
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
}