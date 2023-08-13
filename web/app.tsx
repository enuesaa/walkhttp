import { StrictMode } from 'react'
import { RouterProvider, createBrowserRouter } from 'react-router-dom'
// import './styles/global.css'
import TopPage from './pages/index'

const router = createBrowserRouter([
  {
    path: '/',
    element: <TopPage />,
  },
])

export const App = () => {
  return (
    <>
      <StrictMode>
        <RouterProvider router={router} />
      </StrictMode>
    </>
  )
}