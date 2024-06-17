import Link from 'next/link'
import styles from '@/components/common/Header.css'

export const Header = () => {
  return (
    <header className={styles.main}>
      <Link href='/' className={styles.heading}>
        walkhttp
      </Link>
    </header>
  )
}
