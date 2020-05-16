export interface Repo {
  owner: string;
  repoName: string;
}

export default class TestedRepoModel {
  starNum: number;
  goFiles: FileModel[];
}

export interface TestedFileModel {
  name: string;
  markedContent: MarkLine[];
  testCoverage: CoverageBlockModel[];
}

export interface MarkLine {
  line: string;
  tested: boolean;
}
export interface FileModel {
  name: string;
  content: string;
  testCoverage: CoverageBlockModel[];
}

export interface CoverageBlockModel {
  startLine: number;
  endLine: number;
}
