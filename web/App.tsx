import { createBrowserRouter, RouterProvider } from 'react-router-dom'

// see https://dev.to/franciscomendes10866/file-based-routing-using-vite-and-react-router-3fdo

const pages: Record<string, any> = import.meta.glob('./pages/**/*.tsx', { eager: true })

const routes = []
for (const path of Object.keys(pages)) {
  const Page = pages[path].default
  const normalized = path
    .replace('./pages', '')
    .replace('/index.tsx', '/')
    .replace('.tsx', '')
    .toLowerCase()
  routes.push({
    path: normalized,
    element: <Page />,
  })
}

const router = createBrowserRouter(routes)

export const App = () => {
  return (
    <RouterProvider router={router} />
  )
}
