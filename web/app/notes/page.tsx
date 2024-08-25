"use client";

import { useSearchParams } from "next/navigation";
import { useState } from "react";

type NoteBlocks = {
    blocks: string
}

export default function Note() {
    const searchParams = useSearchParams();
    const code = searchParams.get("code");
    const [note, setNote] = useState<string>()
  
  const getNoteBody = async () => {
    const data: NoteBlocks = await fetch(`http://localhost:3000/api/notes?code=${code}`).then(res => res.json())
    setNote(data.blocks)
    document.getElementById("editor")!.innerHTML = note!
  }

  const copyToClipboard = (v: string) => {
    navigator.clipboard.writeText(v)
  }

  getNoteBody()

  return (
    <div className="w-1/2 my-10 mx-auto">
     <div className="flex justify-between mb-2">
     <label className=" my-auto" htmlFor="code"><span className="font-bold">Code:</span> {code}</label>
      <button onClick={() => copyToClipboard(document.getElementById("editor")!.innerText)} className="text-gray-700 ring-1 ring-inset ring-gray-300 px-2 py-1 bg-white hover:bg-gray-50 rounded-md">
        <svg className="lucide lucide-copy w-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><rect width="14" height="14" x="8" y="8" rx="2" ry="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/></svg>
      </button>
     </div>
      <div id="editor" className="w-full mx-auto rounded shadow-sm bg-white ring-1 ring-inset ring-gray-300 min-h-[500px] p-1 focus:outline-none">
      </div>
    </div>
  );
}
