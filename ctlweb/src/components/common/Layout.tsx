import { ReactNode } from 'react'
import { Header } from './Header'

type Props = {
  children: ReactNode
}
export const Layout = ({ children }: Props) => {
  return (
    <div className='container mx-auto'>
      <Header />
      {children}
    </div>
  )
}
