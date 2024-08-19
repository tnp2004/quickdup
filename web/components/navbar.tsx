import Link from "next/link";

export default function Navbar({name}: {name: string}) {
    return (
        <div className="p-2">
            <Link className="font-bold drop-shadow-md" href={"/"} >{name}</Link>
        </div>
    )
}