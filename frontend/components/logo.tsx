interface LogoProps {
    width?: number
    height?: number
    className?: string
}

export default function GradeFlowLogo({ width = 200, height = 100, className = "" }: LogoProps) {
    return (
        <svg
            width={width}
            height={height}
            viewBox="0 0 200 60"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
            className={className}
        >
            {/* Background circle for icon */}
            <rect x="4" y="8" width="44" height="44" rx="12" fill="#2563EB" />

            {/* Clipboard/Document shape */}
            <rect x="12" y="16" width="28" height="32" rx="4" fill="white" />
            <rect x="18" y="12" width="16" height="4" rx="2" fill="#E5E7EB" />

            {/* Checkmarks */}
            <path d="M16 26L19 29L26 22" stroke="#2563EB" strokeWidth="2.5" strokeLinecap="round" strokeLinejoin="round" />
            <path d="M16 34L19 37L26 30" stroke="#2563EB" strokeWidth="2.5" strokeLinecap="round" strokeLinejoin="round" />

            {/* Lines representing text */}
            <rect x="28" y="25" width="8" height="2" rx="1" fill="#D1D5DB" />
            <rect x="28" y="33" width="8" height="2" rx="1" fill="#D1D5DB" />

            {/* Text "GradeFlow" */}
            <text
                x="60"
                y="25"
                fontFamily="system-ui, -apple-system, sans-serif"
                fontSize="20"
                fontWeight="700"
                fill="#1F2937"
            >
                Grade
            </text>
            <text
                x="60"
                y="42"
                fontFamily="system-ui, -apple-system, sans-serif"
                fontSize="20"
                fontWeight="700"
                fill="#2563EB"
            >
                Flow
            </text>

            {/* Decorative elements */}
            <circle cx="180" cy="20" r="2" fill="#2563EB" opacity="0.3" />
            <circle cx="185" cy="30" r="1.5" fill="#2563EB" opacity="0.5" />
            <circle cx="175" cy="35" r="1" fill="#2563EB" opacity="0.4" />
        </svg>
    )
}

