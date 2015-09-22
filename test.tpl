\documentclass[a4paper]{report}

\usepackage[left=1.25cm,right=1.25cm,top=1cm,bottom=2cm,bindingoffset=0cm]{geometry}
\usepackage[utf8]{inputenc}
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
	escapeinside={\%*}{*)}
}
{{range .Questions}}
	\begin{framed}
		{{if .Code}}
			\begin{minipage}[t]{0.45\textwidth}
			\flushleft
		{{end}}
		{{.Question}} \newline
		\begin{enumerate}
		{{range .Variants}}
			\item {{.}}
		{{end}}
		\end{enumerate}
		{{if .Code}}
			\end{minipage}
			\hfill
			\begin{minipage}[t]{0.50\textwidth}
			\hspace*{0pt}
			\begin{lstlisting}
			{{ .Code }}
			\end{lstlisting}
			\end{minipage}
		{{end}}
	\end{framed}
{{end}}
\end{document}