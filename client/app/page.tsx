import Link from 'next/link'
import { Button } from '@/components/ui/button'

export default function Home() {
  return (
    <section className='py-24'>
      <div className='container'>
        <h1 className='text-3xl font-bold'>Шинохранилеще</h1>
        <Button asChild>
          <Link href='/users'>Пользователи</Link>
        </Button>
      </div>
    </section>
  )
}
