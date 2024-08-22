"use client";

import { useEditor, EditorContent } from "@tiptap/react";
import StarterKit from "@tiptap/starter-kit";
import Toolbar from "./toolbar";
import { useState } from "react";

export default function Editor() {
  const editor = useEditor({
    extensions: [StarterKit.configure({
      heading: {
        levels: [3],
        HTMLAttributes: {
          class: 'heading-3',
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
          class: 'list list-disc font-bold'
        }
      },
      blockquote: {
        HTMLAttributes: {
          class: 'quote'
        }
      }
    })],
    content: "<p>Hello world</p>",
    editorProps: {
        attributes: {
            class: "min-h-[150px] p-1 focus:outline-none"
        }
    }
  });

  const [title, setTitle] = useState("Untitled")

  return (
    <div>
      <div className="w-1/2 mx-auto flex flex-col justify-center">
          <input className="text-3xl font-bold mb-2 focus:outline-none bg-transparent" type="text" value={title} onChange={(e) => setTitle(e.target.value)} />
          <Toolbar editor={editor} />
          <EditorContent className="w-full mx-auto border-2 rounded shadow-sm" editor={editor}/>
      </div>
      <button onClick={() => console.log(editor?.getHTML())}>get html</button>
    </div>
  );
}
