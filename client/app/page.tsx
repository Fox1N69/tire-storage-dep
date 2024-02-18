import Link from "next/link";
import { Button } from "@/components/ui/button";
import "@/app/sass/index.scss";

export default function Home() {
  return (
    <section className="py-24">
      <div className="container ">
        <h1 className="text-3xl font-bold text-center">
          Система для склада шин
        </h1>
        <div className="buttons min-w-full items-center flex  gap-5 container mx-0 justify-center mt-10">
          <Button asChild className="w-34">
            <Link href="/users">Клиентская база</Link>
          </Button>
          <Button asChild className=" w-34">
            <Link href="/storagerequests">Работа с заявками</Link>
          </Button>
        </div>
      </div>
    </section>
  );
}
