import * as fromRepo from './repo.actions';

describe('loadRepos', () => {
  it('should return an action', () => {
    expect(fromRepo.loadRepos().type).toBe('[Repo] Load Repos');
  });
});
