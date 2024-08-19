import Link from "next/link";

export default function Home() {
  return (
    <div>
      <div className="mx-auto w-fit flex flex-col justify-center gap-y-10 my-60">
        <p className="text-4xl text-sky-400 drop-shadow-md w-3/5 text-center font-bold mx-auto">
          QuickNote is your go-to tool for sharing notes easy and swiftly with others.
        </p>
        <p className="text-xl w-[45%] px-20 mx-auto text-center">
          No more clutter—just create, share, and watch your notes self-expire
          in few minutes. Perfect for quick thoughts, meeting recaps, or any
          information you need to pass along.
        </p>
        <Link href={"/note/create"} className="border-2 border-sky-200 shadow-xl w-fit px-6 py-2 mx-auto font-bold rounded-lg bg-sky-400 hover:bg-sky-400/80 text-slate-50">
          take a note
        </Link>
      </div>
    </div>
  );
}
