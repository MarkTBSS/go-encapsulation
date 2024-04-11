Encapsulation คือความสามารถในการซ่อน data ที่มีอยู่กับคลาสนั้นๆ ป้องกันการเข้าถึง data จาก 3rd party และ Class อื่นๆ, ในส่วนของการแก้ไขค่า หรือเข้าถึงค่า ต้องทำผ่าน method เข้านั้น

Encapsulation คือ คุณสมบัติในการเขียนโปรแกรมเชิงวัตถุแล้วมีการกำหนดการเข้าสมาชิกภายใน Class ไม่ว่าภายนอกหรือภายในก็ตามจะถูกนำไปใช้เพื่อป้องกันข้อมูลภายในให้มีความปลอดภัยและเป็นความลับ

ทำไมเราถึงต้องหุ้มเอาไว้ ห่อแล้วจะได้ประโยชน์อะไร ลองนึกภาพถึงโปรแกรมที่มี object อยู่หลายตัวซึ่งแต่ละตัวก็มีการติดต่อรับส่งข้อมูลกัน ถ้าไม่มีการห่อหุ้มเอาไว้ แต่ละ object เข้าถึงและเปลี่ยนแปลงคุณสมบัติของ object อื่นได้ทั้งหมด อาจจะทำให้เกิดความผิดพลาดในการทำงานได้ถ้าถูกเปลี่ยนแปลงคุณสมบัติโดยไม่มีการควบคุม