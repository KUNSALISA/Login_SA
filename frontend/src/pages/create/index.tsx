import React, { useState } from "react";
import { Form, Input, Button, message } from "antd";
import { useNavigate } from "react-router-dom";
import { CreateUser } from "../../services/https";

function Login() {
  const navigate = useNavigate();
  const [messageApi, contextHolder] = message.useMessage();

  const onFinish = async (values: { Email: string; password: string }) => {
    let res = await CreateUser(values); // เปลี่ยน CreateUser เป็น API สำหรับการ login
    if (res) {
      messageApi.open({
        type: "success",
        content: "เข้าสู่ระบบสำเร็จ",
      });
      setTimeout(function () {
        navigate("/dashboard"); // เปลี่ยนไปหน้าหลังจาก login สำเร็จ
      }, 2000);
    } else {
      messageApi.open({
        type: "error",
        content: "เกิดข้อผิดพลาดในการเข้าสู่ระบบ!",
      });
    }
  };

  return (
    <div>
      {contextHolder}
      <h2>เข้าสู่ระบบ</h2>
      <Form name="login" layout="vertical" onFinish={onFinish}>
        <Form.Item
          label="อีเมล"
          name="Email"
          rules={[
            { type: "email", message: "รูปแบบอีเมลไม่ถูกต้อง!" },
            { required: true, message: "กรุณากรอกอีเมล!" },
          ]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label="รหัสผ่าน"
          name="password"
          rules={[
            { required: true, message: "กรุณากรอกรหัสผ่าน!" },
          ]}
        >
          <Input.Password />
        </Form.Item>

        <Form.Item>
          <Button type="primary" htmlType="submit">
            เข้าสู่ระบบ
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
}

export default Login;
