import java.util.Scanner;

public class App {
    public static void main(String[] args) {
        // Tạo một đối tượng Scanner để đọc dữ liệu từ đầu vào
        Scanner scanner = new Scanner(System.in);

        // In ra câu hỏi
        System.out.print("Tên bạn là gì? ");

        // Đọc tên người dùng nhập vào
        String name = scanner.nextLine();

        // In ra lời chào
        System.out.println("Chào bạn " + name);

        // Đóng Scanner để giải phóng tài nguyên
        scanner.close();
    }
}
