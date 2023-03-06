// считайте что fmt уже импортирован и main объявлен
func test(x1 *int, x2 *int) {
	// здесь пишите ваш код
	*x1 = *x1 * *x2
	fmt.Println(*x1)
}