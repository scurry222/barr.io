import React from "react";


// VERY unsafe login method. Change immediately after backend is set up successfully

function Login() {
  return (
    <div >
      <h1>LOG IN</h1>
      <form className="flex flex-col p-24 margin-auto">
        <label>Name: </label>
        <input className="border-black border-2 border-solid" type="text" id="user-name" name="user-name"/>
        <br />
        <label>Email: </label>
        <input className="border-black border-2 border-solid"  type="text" id="email" name="email"/>
        <br />
        <label>Password: </label>
        <input className="border-black border-2 border-solid" type="text" id="password" name="password"/>
        <br />
        <label>Re-enter password: </label>
        <input className="border-black border-2 border-solid" type="text" id="password" name="password"/>
        <input type="submit" value="Submit" />
      </form>
    </div>
  )
}

export default Login;