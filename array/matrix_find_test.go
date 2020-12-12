package array

import (
	"fmt"
	"testing"
)

func TestMatrixFind(t *testing.T) {
	matrix := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15}}

	fmt.Println(MatrixFind(matrix, 7))
	fmt.Println(MatrixFind(matrix, 1))
	fmt.Println(MatrixFind(matrix, 11))
	fmt.Println(MatrixFind(matrix, 16))
	fmt.Println("تم حظرك من دخول جميع غرف الدردشة إلى الأبد بسبب سلوكك المخالف للقواعد. إذا كنت تشعر أن قرارنا قد تم عن طريق الخطأ ، فالرجاء إرسال رسالة بريد إلكتروني إلى العنوان moderation_user@funshareapp.com والأيدي الخاص بك ولماذا تعتقد أنه يتعين علينا إلغاء حظر حسابك. سنعود إليك في خلال 3 أيام عمل بعد تلقي البريد الإلكتروني والمعلومات الخاصة بك.")

}
