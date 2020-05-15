export interface Repo {
  owner: string;
  repoName: string;
}

export default class TestedRepoModel {
  starNum: number;
  goFiles: FileModel[];
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
