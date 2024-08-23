"use client";

import { useEditor, EditorContent } from "@tiptap/react";
import StarterKit from "@tiptap/starter-kit";
import Toolbar from "./toolbar";
import { useState } from "react";

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
    content: "<p>Hello world</p>",
    editorProps: {
        attributes: {
            class: "min-h-[500px] p-1 focus:outline-none"
        }
    },
    immediatelyRender: false
  });

  const [title, setTitle] = useState("Untitled")
  
  return (
    <div>
      <div className="w-1/2 mx-auto flex flex-col justify-center">
          <input className="text-3xl font-bold mb-2 focus:outline-none bg-transparent" type="text" value={title} onChange={(e) => setTitle(e.target.value)} />
          <Toolbar editor={editor} />
          <EditorContent className="w-full mx-auto rounded shadow-sm bg-white ring-1 ring-inset ring-gray-300" editor={editor}/>
      </div>
      <button onClick={() => console.log(editor?.getHTML())}>get html</button>
    </div>
  );
}
