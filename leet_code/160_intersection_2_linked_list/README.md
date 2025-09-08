Bài Intersection of Two Linked Lists là một trong những câu LeetCode kinh điển, kiểm tra khả năng thao tác với linked list và tư duy con trỏ. Dưới đây là phân tích chi tiết và cách giải.

⸻

💡 Bài toán:

Cho hai singly linked list headA và headB, hãy tìm node mà tại đó hai list giao nhau (nếu có). Trả về node đó (pointer), hoặc nil nếu không giao.

Constraints:
•	Không được sửa đổi cấu trúc của list.
•	Thời gian: O(n)
•	Không dùng thêm bộ nhớ (O(1) space)

⸻

📌 Ý tưởng giải:

✅ Cách tối ưu nhất (Two pointers):

Dùng hai con trỏ a và b, bắt đầu từ headA và headB. Mỗi khi một con trỏ đi đến cuối list, ta cho nó đi lại từ đầu của list còn lại. Chúng sẽ gặp nhau tại node giao nhau (hoặc nil nếu không giao).

💥 Tại sao hiệu quả?

Giả sử:
•	Độ dài list A = a + c
•	Độ dài list B = b + c

(trong đó c là đoạn giao nhau, a và b là phần không chung)

Khi ta cho hai con trỏ đi theo:
•	ptrA: A → B
•	ptrB: B → A

Sau a + b + c bước, cả hai con trỏ sẽ đến cùng 1 node nếu có intersection.

⸻

✅ Code Go:

func getIntersectionNode(headA, headB *ListNode) *ListNode {
if headA == nil || headB == nil {
return nil
}

    a, b := headA, headB

    for a != b {
        if a != nil {
            a = a.Next
        } else {
            a = headB
        }

        if b != nil {
            b = b.Next
        } else {
            b = headA
        }
    }

    return a // hoặc b, vì a == b
}



⸻

🧠 Edge Cases:
•	Một trong hai list là nil → return nil
•	Không giao nhau → cả hai con trỏ cùng về nil
•	Giao nhau tại đầu hoặc cuối → vẫn đúng

⸻

Nếu bạn cần mình vẽ sơ đồ minh họa hoặc viết test case cụ thể bằng Go, mình sẵn sàng nhé!