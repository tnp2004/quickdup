import { Editor } from "@tiptap/react";
import HeadingDropdown from "./headingDropdown";
import { isActive } from "@/utils";

export default function Toolbar({ editor }: {editor: Editor | null}) {
    return (
      <div className="flex gap-2 mb-1 py-1">
        <button className={`${isActive(editor,"bold")} toolBtn`}
          onClick={() => editor?.chain().focus().toggleBold().run()}>bold</button>

        <HeadingDropdown editor={editor}/>
        
          <button className={`${isActive(editor,"italic")} toolBtn`}
          onClick={() => editor?.chain().focus().toggleItalic().run()}>italic</button>
        <button className={`${isActive(editor,"list")} toolBtn`}
          onClick={() => editor?.chain().focus().toggleBulletList().run()}>list</button>
      </div>
    )
}