import { Link } from 'react-router-dom'

//TSX element for sidebar link component
const SidebarLink = ({
    to,
    children,
}: {
    to: string
    children: React.ReactNode
}) => {
    return (
        <li className="mb-2 ">
            <Link to={to}>{children}</Link>
        </li>
    )
}

const Sidebar = () => {
    return (
        <div className="flex flex-col bg-gray-200 w-64">
            <ul className="p-4">
                <SidebarLink to="/dashboard/">Home</SidebarLink>
                <SidebarLink to="/dashboard/transactions">
                    Transactions
                </SidebarLink>
                <SidebarLink to="/dashboard/categories">Categories</SidebarLink>
                <SidebarLink to="/dashboard/budgets">Budgets</SidebarLink>
            </ul>
        </div>
    )
}

export default Sidebar
