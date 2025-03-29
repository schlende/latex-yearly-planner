\vskip 5mm
\textbf{Goal Name} \underline{\hspace{5cm}}
\vskip 5mm

\textbf{Deadline:} \underline{\hspace{5cm}}
\vskip 5mm

\textbf{Goal Description:}
\myMash{3}{\myNumDotWidthFull}
\vskip 5mm

\textbf{What makes this goal important to you?}
\myMash{3}{\myNumDotWidthFull}
\vskip 5mm

\textbf{Goal partners:}
\vskip 2mm
\begin{tabularx}{\linewidth}{|X|X|}
  \hline
  \myLineHeightButLine{} & \myLineHeightButLine{} \\
  \hline
  \myLineHeightButLine{} & \myLineHeightButLine{} \\
  \hline
\end{tabularx}
\vskip 5mm

\textbf{Milestones:}
\vskip 2mm
\begin{tabularx}{\linewidth}{|X|c|}
  \hline
  \textbf{Milestone} & \textbf{Date} \\
  \hline
  \myLineHeightButLine{} & \\
  \hline
  \myLineHeightButLine{} & \\
  \hline
  \myLineHeightButLine{} & \\
  \hline
  \myLineHeightButLine{} & \\
  \hline
\end{tabularx}
\vskip 5mm

\textbf{Daily sacrifices:}
\vskip 2mm
\begin{tabularx}{\linewidth}{|X|}
  \hline
  \myLineHeightButLine{} \\
  \hline
  \myLineHeightButLine{} \\
  \hline
  \myLineHeightButLine{} \\
  \hline
\end{tabularx}
\vskip 5mm

\textbf{Repetition Record:}
\vskip 2mm
\begin{tabularx}{\linewidth}{|c|c|c|c|c|c|c|}
  \hline
  \textbf{Mon} & \textbf{Tue} & \textbf{Wed} & \textbf{Thu} & \textbf{Fri} & \textbf{Sat} & \textbf{Sun} \\
  \hline
  & & & & & & \\
  \hline
  & & & & & & \\
  \hline
  & & & & & & \\
  \hline
  & & & & & & \\
  \hline
\end{tabularx}

{{ if $.Cfg.Dotted -}} 
\vskip\myLenLineHeightButLine\myMash{\myNumDotHeightFull}{\myNumDotWidthFull} 
{{- else -}}
\vbox to \dimexpr\textheight-\pagetotal-\myLenLineHeightButLine\relax {%
  \leaders\hbox to \linewidth{\textcolor{\myColorGray}{\rule{0pt}{\myLenLineHeightButLine}\hrulefill}}\vfil
}
{{end}}
