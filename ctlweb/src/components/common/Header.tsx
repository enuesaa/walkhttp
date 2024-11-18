import Link from 'next/link'
import { MakeFormDialog } from './MakeFormDialog'

export const Header = () => {
  return (
    <header className='flex'>
      <div className='flex-auto'>
        <Link href='/' className='inline-block text-lg font-bold px-3 py-2 hover:bg-gray rounded'>
          walkhttp
        </Link>
      </div>

      <div className='w-20 text-right'>
        <MakeFormDialog />
      </div>
    </header>
  )
}
