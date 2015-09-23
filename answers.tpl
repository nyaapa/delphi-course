\documentclass[a4paper]{report}

\usepackage[left=1.25cm,right=1.25cm,top=1cm,bottom=2cm,bindingoffset=0cm]{geometry}
\usepackage[utf8x]{inputenc}
\usepackage{ucs}
\usepackage[russian]{babel}

\usepackage{framed}
\usepackage{etoolbox}
\usepackage{color}
\usepackage{listings}
\BeforeBeginEnvironment{framed}{\begin{minipage}{\linewidth}}
\AfterEndEnvironment{framed}{\end{minipage}\par}
\begin{document}
\pagestyle{empty}
\lstset{
	language=Pascal,
	basicstyle=\small\ttfamily,
	numbers=left,
	numberstyle=\tiny,
	stepnumber=1,
	numbersep=5pt,
	backgroundcolor=\color{white},
	showspaces=false,
	showstringspaces=false,
	showtabs=false,
	frame=single,
	tabsize=2,
	captionpos=t,
	breaklines=true,
	breakatwhitespace=false,
	escapeinside={\%*}{*)},
	inputencoding=utf8x,
	extendedchars=\true
}
\begin{huge} \textbf{ {{ .Title }} } \end{huge}

{{range .Questions}}
	{{.No}}: {{.Answers}} \newline
{{end}}
\end{document}