import Link from 'next/link'
import { MakeFormDialog } from './MakeFormDialog'

export const Header = () => {
  return (
    <header className='p-2'>
      <Link href='/' className='text-gray-100 no-underline h-12 leading-[50px] text-lg font-bold m-2'>
        walkhttp
      </Link>

      <MakeFormDialog />
    </header>
  )
}
