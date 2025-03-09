import Link from 'next/link'
import { MakeFormDialog } from './MakeFormDialog'
import { ExportButton } from './ExportButton'

export const Header = () => {
  return (
    <header className='flex'>
      <div className='flex-auto'>
        <Link href='/' className='inline-block text-lg font-bold px-3 py-2 hover:bg-stone-700'>
          walkhttp
        </Link>
      </div>

      <div className='w-32 text-right'>
        <MakeFormDialog />
        <ExportButton />
      </div>
    </header>
  )
}
