'use client'
import "@/app/globals.css";


import { ChangeEventHandler, SetStateAction, useState } from "react";

export default function TextBox() {
  const [text, setText] = useState("");

  return (
    <div className="textbox">
      <form>
        <input
          className="border border-2 border-black"
          type="text"
          value={text}
          onChange={(e) => setText(e.target.value)} />
        <button onClick={(e) => e.preventDefault()}>{"->"}</button>
      </form>
    </div>
  )
}