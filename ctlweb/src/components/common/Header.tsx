import Link from 'next/link'

export const Header = () => {
  return (
    <header className='h-12 leading-[50px] text-lg font-bold'>
      <Link href='/' className='text-gray-100 mx-2 no-underline'>
        walkhttp
      </Link>
    </header>
  )
}
