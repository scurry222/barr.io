import Link from "next/link";

type Props = {
  pageName: string,
  title: string
}

function NavigateToPage({ pageName, title }: Props) {
  return (
    <Link href={pageName}>{title}</Link>
  )
}

export default NavigateToPage;