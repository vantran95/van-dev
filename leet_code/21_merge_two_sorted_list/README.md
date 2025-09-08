📘 Mục tiêu:

Hàm mergeTwoLists nhận vào 2 linked list đã được sắp xếp tăng dần, trả về 1 linked list cũng sắp xếp tăng dần, chứa tất cả các phần tử của cả hai list.

⸻

🧩 Phân tích từng dòng:

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	•	Định nghĩa hàm với 2 tham số là list1 và list2 kiểu *ListNode.

⸻



    dummy := &ListNode{} // node giả để dễ xử lý
    current := dummy

	•	Tạo một node giả dummy với giá trị mặc định (0).
	•	Dùng con trỏ current để duyệt và xây danh sách kết quả.
	•	Mục đích của dummy là giúp không cần xử lý riêng trường hợp đầu list, vì dummy.Next sẽ là head của danh sách kết quả.

⸻



    for list1 != nil && list2 != nil {

	•	Lặp lại khi cả list1 và list2 đều còn phần tử.

⸻



        if list1.Val < list2.Val {
            current.Next = list1
            list1 = list1.Next
        } else {
            current.Next = list2
            list2 = list2.Next
        }
        current = current.Next

	•	So sánh Val của 2 node đầu:
	•	Node nhỏ hơn được nối vào current.Next
	•	Con trỏ tương ứng được di chuyển sang node tiếp theo.
	•	Cập nhật current để trỏ đến node mới được thêm.

⸻



    if list1 != nil {
        current.Next = list1
    } else {
        current.Next = list2
    }

	•	Sau vòng lặp, một trong hai list có thể vẫn còn phần tử (list còn lại đã hết).
	•	Gắn trực tiếp phần còn lại vào danh sách kết quả vì nó đã được sắp xếp sẵn.

⸻



    return dummy.Next
}

	•	Trả về list kết quả bắt đầu từ dummy.Next (bỏ qua node giả dummy).

⸻

🧠 Hình dung trực quan:

Giả sử:

list1: 1 → 3 → 5  
list2: 2 → 4 → 6

Sau khi chạy hàm, bạn sẽ nhận được:

dummy → 1 → 2 → 3 → 4 → 5 → 6

Trả về dummy.Next chính là:

1 → 2 → 3 → 4 → 5 → 6



⸻

✅ Ưu điểm:

•	Dễ hiểu, gọn gàng, không cần xử lý đặc biệt cho phần đầu list.

•	O(n + m) thời gian, O(1) bộ nhớ (không dùng mảng hay slice phụ).

⸻