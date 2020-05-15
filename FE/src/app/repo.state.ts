import TestedRepoModel from './repo.model';

export default class RepoState {
  loading: boolean;
  loaded: boolean;
  error: string;
  tested: TestedRepoModel;
}

export const initializeState = (): RepoState => {
  return {
    loading: false,
    loaded: false,
    error: undefined,
    tested: undefined,
  };
};
