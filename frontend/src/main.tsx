import React from 'react'
import ReactDOM from 'react-dom/client'
import Root from './routes/root'
import './index.css'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import Error404Page from './components/Error404Page'
import Dashboard from './routes/dashboard'
import { createTheme, CssBaseline, ThemeProvider } from '@mui/material'

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
        path: '/register',
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

const darkTheme = createTheme({
    palette: {
        mode: 'dark',
    },
})

ReactDOM.createRoot(document.getElementById('root')!).render(
    <React.StrictMode>
        <ThemeProvider theme={darkTheme}>
            <RouterProvider router={router} />
            <CssBaseline />
        </ThemeProvider>
    </React.StrictMode>
)
