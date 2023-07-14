import React from 'react'
import ReactDOM from 'react-dom/client'
import Root from './routes/root'
import './index.css'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import Error404Page from './components/Error404Page'
import Dashboard from './routes/dashboard'
import Signup from './routes/signup'

const router = createBrowserRouter([
    {
        path: '/',
        element: <Root />,
        errorElement: <Error404Page />,
    },
    {
        path: '/login',
        errorElement: <Error404Page />,
    },
    {
        path: '/signup',
        element: <Signup />,
        errorElement: <Error404Page />,
    },
    {
        path: '/dashboard',
        element: <Dashboard />,
        errorElement: <Error404Page />,
        children: [
            {
                path: 'dashboard/home',
                element: <div>Dashboard</div>,
            },
        ],
    },
])

ReactDOM.createRoot(document.getElementById('root')!).render(
    <React.StrictMode>
        <RouterProvider router={router} />
    </React.StrictMode>
)
