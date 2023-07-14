import { useState } from 'react'

const Signup = () => {
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')

    const handleSubmit: React.FormEventHandler<HTMLFormElement> = (e) => {
        e.preventDefault()
        console.log(email)
        console.log(password)
    }

    const validPassword = () => {
        return password.length > 8
    }

    return (
        <form className="flex flex-col items-center justify-center h-screen">
            <h1 className="text-5xl font-bold">Sign Up</h1>
            <div className="form-control w-full max-w-xs">
                <label className="label">
                    <span className="label-text">Email</span>
                </label>
                <input
                    type="email"
                    placeholder="YourEmail@Here.com"
                    className="input input-bordered input-primary w-full max-w-xs"
                    onChange={(e) => setEmail(e.target.value)}
                />
            </div>
            <div className="form-control w-full max-w-xs">
                <label className="label">
                    <span className="label-text">Password</span>
                </label>
                <input
                    type="password"
                    placeholder=""
                    className="input input-bordered input-primary w-full max-w-xs"
                    onChange={(e) => setPassword(e.target.value)}
                />
                <span className="label-text">
                    <ul className="">
                        <li
                            className={`${
                                validPassword() ? 'text-success' : 'text-error'
                            }`}
                        >
                            Minimum Length of 8
                        </li>
                        <li
                            className={`${
                                password.match(/[0-9]/g)
                                    ? 'text-success'
                                    : 'text-error'
                            }`}
                        >
                            Contains Number or Special Character
                        </li>
                        <li
                            className={`${
                                password.match(/[a-z]/g)
                                    ? 'text-success'
                                    : 'text-error'
                            }`}
                        >
                            Uppercase Character
                        </li>
                    </ul>
                </span>
            </div>

            <button
                type="submit"
                className="btn btn-primary btn-block max-w-xs"
                disabled={!validPassword()}
                onClick={handleSubmit}
            >
                Sign Up
            </button>
        </form>
    )
}

export default Signup
