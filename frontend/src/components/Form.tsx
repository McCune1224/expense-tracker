import React from 'react'

const Form = () => {
    return <form>
        <label htmlFor="email">Email</label>
        <input type="email" id="email" name="email" />
        <label htmlFor="password">Password</label>

        <input type="password" id="password" name="password" />
        <button type="submit">Submit</button>
    </form>

}

export default Form
