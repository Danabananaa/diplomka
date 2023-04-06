// ID          int64      `json:"id"`
// UserID      int64      `json:"user_id"`
// TypeID      int64      `json:"type_id"`
// Description string     `json:"description"`
// Amount      int64      `json:"amount"`
// Date        CustomTime `json:"date"`
// Status      bool       `json:"status"`

import { getCurrentDate } from "../API";

export const sendDebt = async (status, typeID, amount, description, navigate) => {
    const date = getCurrentDate();
    const token = localStorage.getItem("token");
    const data = {
      type_id: parseInt(typeID, 10),
      description: description,
      amount: parseInt(amount, 10),
      date: date,
      status: status === "Debt" ? true : false,
    };
    console.log(data.income_type_id);
    await fetch(`/debt`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `${token}`,
      },
      body: JSON.stringify(data),
    }).then(async (response) => {
      if (!response.ok) {
        console.log("something went wrong");
      } else {
        console.log("Income sent");
      }
    });
  };