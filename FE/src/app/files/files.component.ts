import { Component, OnInit } from '@angular/core';
import { Store, select, resultMemoize } from '@ngrx/store';
import { Observable } from 'rxjs';
import { tap } from 'rxjs/operators';
import { ActivatedRoute } from '@angular/router';

import { Repo, FileModel, TestedFileModel, MarkLine } from '../repo.model';
import RepoState from '../repo.state';
import * as RepoActions from '../repo.action';
import TestedRepoModel from '../repo.model';

@Component({
  selector: 'files',
  templateUrl: './files.component.html',
  styleUrls: ['./files.component.scss'],
})
export class FilesComponent implements OnInit {
  path: any;
  repoSub$: Observable<RepoState>;
  repoObject: RepoState;
  repo: TestedRepoModel;
  markTestedFiles: TestedFileModel[];
  currentFile: TestedFileModel;
  displayedColumns: string[] = ['name'];
  starNum: string;

  constructor(
    private store: Store<{ repo: RepoState }>,
    private route: ActivatedRoute
  ) {
    this.repoSub$ = this.store.select('repo').pipe(
      tap((result) => {
        this.repoObject = result;
        this.repo = this.repoObject.tested;
        if (this.repo) {
          this.markTestedFiles = this.repo.goFiles.map((x) => {
            let splited = x.content.split('\n');
            let markedContent = splited.map((line) => {
              let markedLines: MarkLine = {
                line: line,
                tested: false,
              };
              return markedLines;
            });

            let testedFile: TestedFileModel = {
              name: x.name,
              markedContent: markedContent,
              testCoverage: x.testCoverage,
            };
            return testedFile;
          });
          this.currentFile = this.markTestedFiles[0];
          this.starNum = (this.repo.starNum / 1000).toFixed(1) + ' k';
        }
      })
    );
  }

  changeFile(file) {
    this.currentFile = file;
  }

  runTest() {
    let newMarkTestedFiles = this.markTestedFiles.map((file) => {
      let testedLines = new Set();
      if (file.testCoverage) {
        file.testCoverage.forEach((x) => {
          for (let i = x.startLine; i <= x.endLine; i++) {
            testedLines.add(i);
          }
        });
      }

      let markedLines: MarkLine[] = [];

      for (let i = 0; i < file.markedContent.length; i++) {
        let markLine: MarkLine = {
          line: file.markedContent[i].line,
          tested: false,
        };

        if (testedLines.has(i + 1)) {
          markLine.tested = true;
        } else {
          markLine.tested = false;
        }
        markedLines.push(markLine);
      }

      let testedFile: TestedFileModel = {
        name: file.name,
        markedContent: markedLines,
        testCoverage: file.testCoverage,
      };

      return testedFile;
    });

    console.log(newMarkTestedFiles);
    this.markTestedFiles = newMarkTestedFiles;
    console.log(this.markTestedFiles);
  }

  ngOnInit(): void {
    this.path = this.route.snapshot.paramMap.get('path');
    console.log(this.path);
  }
}
