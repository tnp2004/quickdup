'use client'

import Link from "next/link";
import { redirect } from "next/navigation";
import { useState } from "react";

export default function Navbar({name}: {name: string}) {
    const [code, setCode] = useState<string>("")
    return (
        <div className="flex justify-between p-2">
            <Link className="font-bold drop-shadow-md" href={"/"} >{name}</Link>
            <div className="flex gap-1">
                <input autoComplete="off" onChange={(e) => setCode(e.target.value)} value={code} id="codeInput" className="h-7 px-2 drop-shadow-sm rounded text-center ring-1 ring-inset ring-gray-300" placeholder="enter code" type="text" />
                <Link href={`/notes?code=${code}`} >
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="lucide lucide-search w-5 text-gray-600 my-auto"><circle cx="11" cy="11" r="8"/><path d="m21 21-4.3-4.3"/></svg>
                </Link>
            </div>
            <div></div>
        </div>
    )
}