"use client";

import { useEditor, EditorContent } from "@tiptap/react";
import StarterKit from "@tiptap/starter-kit";
import Toolbar from "./toolbar";
import { useState } from "react";
import Link from "next/link";

export default function Editor() {
  const editor = useEditor({
    autofocus: true,
    extensions: [StarterKit.configure({
      heading: {
        levels: [1,2,3],
        HTMLAttributes: {
          class: 'heading',
        },
      },
      bold: {
        HTMLAttributes: {
          class: 'bold'
        }
      },
      italic: {
        HTMLAttributes: {
          class: 'italic'
        }
      },
      bulletList: {
        HTMLAttributes: {
          class: 'bulletList'
        }
      }
    })],
    content: "",
    editorProps: {
        attributes: {
            class: "min-h-[500px] p-1 focus:outline-none"
        }
    },
    immediatelyRender: false
  });

  const [code, setCode] = useState<string>("")
  const shareNote = async () => {
    const data = await(await fetch(`http://localhost:3000/api/notes/create`, {
      method: "POST",
      body: JSON.stringify({
        userID: "",
        blocks: editor?.getHTML()
      })
    })).json()
    setCode(data.code)
  }

  const copyToClipboard = () => {
    navigator.clipboard.writeText(code)
  }
  
  return (
    <div>
      <div className="w-1/2 my-10 mx-auto flex flex-col justify-center">
          <div className="flex gap-2 justify-end">

            <div className={`${code.length == 0 ? "hidden" : "block"} flex gap-1 my-auto`}>
           <span className="my-auto bg-gray-200 px-1 rounded-sm">{code}</span>
            <button onClick={copyToClipboard} className="text-gray-700 ring-1 ring-inset ring-gray-300 px-2 py-1 bg-white hover:bg-gray-50 rounded-md">
              <svg className="lucide lucide-copy w-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><rect width="14" height="14" x="8" y="8" rx="2" ry="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/></svg>
            </button>
              <Link href={`/notes?code=${code}`} className="text-gray-700 ring-1 ring-inset ring-gray-300 px-2 py-1 bg-white hover:bg-gray-50 rounded-md">
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="lucide lucide-link w-4"><path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"/><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"/></svg>
            </Link>
              </div>
              
            <button className="text-sm ring-1 ring-inset ring-gray-300 bg-white hover:bg-gray-50 rounded-md shadow-sm px-3 py-2 h-fit" onClick={shareNote}>share</button>

         </div>
          <Toolbar editor={editor} />
          <EditorContent className="w-full mx-auto rounded shadow-sm bg-white ring-1 ring-inset ring-gray-300" editor={editor}/>
      </div>
    </div>
  );
}
