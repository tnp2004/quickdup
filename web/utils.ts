import { Editor } from "@tiptap/react"

export const isActive = (editor: Editor | null,className: string): string => {
    return editor?.isActive(className) ? "font-bold" : ""
}