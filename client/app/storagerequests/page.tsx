import { User, columns } from './columns'
import { DataTable } from '@/components/data-table'

async function getUsers(): Promise<User[]> {
  const res = await fetch(
    'http://127.0.0.1:4000/api/user'
  )
  const data = await res.json()
  return data
}

export default async function Page() {
  const data = await getUsers()

  return (
    <section className='py-24'>
      <div className='container'>
        <h1 className='mb-6 text-3xl font-bold'>Все заявки</h1>
        <DataTable columns={columns} data={data}/>
      </div>
    </section>
  )
}
