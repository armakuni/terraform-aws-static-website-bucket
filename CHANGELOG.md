# Changelog
All notable changes to this project will be documented in this file. See [conventional commits](https://www.conventionalcommits.org/) for commit guidelines.

- - -
## [v0.4.2](https://github.com/armakuni/terraform-aws-static-website-bucket/compare/v0.4.1..v0.4.2) - 2023-09-25
#### Bug Fixes
- Generate terraform docs - ([1715fcf](https://github.com/armakuni/terraform-aws-static-website-bucket/commit/1715fcf48a063aa622e39026dc0a8d2d6466a416)) - Tom Oram
- Set required terraform version - ([9ea32c8](https://github.com/armakuni/terraform-aws-static-website-bucket/commit/9ea32c8c1e3fa9577a73db96922f1fe1a4460a3f)) - Tom Oram
#### Continuous Integration
- split unit and integrations as separate jobs - ([3da91c8](https://github.com/armakuni/terraform-aws-static-website-bucket/commit/3da91c8e73a5ed0637cd2cef2f5a19417f3cffb4)) - abhisheksr01
#### Documentation
- add pipeline status to README - ([e45c5b7](https://github.com/armakuni/terraform-aws-static-website-bucket/commit/e45c5b710dc319f32571345d01a1cbdd4d7f2e75)) - abhisheksr01
#### Miscellaneous Chores
- **(unit)** fix failing test - ([7d19fd5](https://github.com/armakuni/terraform-aws-static-website-bucket/commit/7d19fd51d3cb6e5e0c9050a361f7acc744effd6f)) - abhisheksr01
#### Refactoring
- **(unit)** use the go-terratest-helper library - ([88e1d64](https://github.com/armakuni/terraform-aws-static-website-bucket/commit/88e1d64581260e47facc3fab6984aff824f8424e)) - abhisheksr01
#### Style
- Add *.go to .editorconfig - ([a2742fa](https://github.com/armakuni/terraform-aws-static-website-bucket/commit/a2742fa9c5def36b6068422aa07cba2ba3e94f28)) - Tom Oram
#### Tests
- **(integration)** refactored folder structure - ([1870e3b](https://github.com/armakuni/terraform-aws-static-website-bucket/commit/1870e3b0b4842ae10b01298da46d122ed6fb4cc5)) - d3vadv3ntur3s
- **(unit)** 3 tests for bucket versioning, acl, and ownership controls - ([49c107c](https://github.com/armakuni/terraform-aws-static-website-bucket/commit/49c107c132aa1002def4eb3ce3cdf26309b8bf25)) - d3vadv3ntur3s
- **(unit)** bucket name and acl from plan output leveraging helper functions - ([bdd158a](https://github.com/armakuni/terraform-aws-static-website-bucket/commit/bdd158a3e0b91eed2b48f7ca36248af64183862e)) - d3vadv3ntur3s
- remove planfilepath reference from the test - ([922dc33](https://github.com/armakuni/terraform-aws-static-website-bucket/commit/922dc33aa818bee5ccbd51021a2b4db96c912bbf)) - abhisheksr01

- - -

## [v0.4.1](https://github.com/armakuni/terraform-aws-static-website-bucket/compare/v0.4.0..v0.4.1) - 2023-09-22
#### Bug Fixes
- Pass grant and owner through to module - ([9bb1d68](https://github.com/armakuni/terraform-aws-static-website-bucket/commit/9bb1d688db790d9c052216bd3385c2e9738b9471)) - Tom Oram
- correct index_document - ([e21f859](https://github.com/armakuni/terraform-aws-static-website-bucket/commit/e21f859e6b050ebf3a3cf5d9de03b404a8d48a8c)) - Tom Oram
#### Continuous Integration
- Use assume-aws-oidc-role action - ([fa50e00](https://github.com/armakuni/terraform-aws-static-website-bucket/commit/fa50e00cc074300ce2fa8ffca2b3ffa9b9361e29)) - Tom Oram

- - -

## [v0.4.0](https://github.com/armakuni/terraform-aws-static-website-bucket/compare/v0.3.1..v0.4.0) - 2023-09-22
#### Continuous Integration
- Make pipeline run on pull request - ([0a5c769](https://github.com/armakuni/terraform-aws-static-website-bucket/commit/0a5c7697c735608aa1d492444fc5ca30d4d64f0e)) - Tom Oram
#### Features
- Make bucket publicly accessible - ([d0655a0](https://github.com/armakuni/terraform-aws-static-website-bucket/commit/d0655a0da971d9e7697adf98642c4dc5ba44d333)) - Tom Oram
#### Tests
- Tidy integration test - ([7c87308](https://github.com/armakuni/terraform-aws-static-website-bucket/commit/7c8730850fc465c70317ea7c2ccfbd1d01ff0200)) - Tom Oram

- - -

## [v0.3.1](https://github.com/armakuni/terraform-aws-static-website-bucket/compare/v0.3.0..v0.3.1) - 2023-09-19
#### Bug Fixes
- update tag-and-release to 0.4.1 - ([c58bc15](https://github.com/armakuni/terraform-aws-static-website-bucket/commit/c58bc1536b4efe97804759d5e8ab034d8b231962)) - Ben Nagy
#### Continuous Integration
- Don't run tests on PR - ([2d8d8fe](https://github.com/armakuni/terraform-aws-static-website-bucket/commit/2d8d8fea4b60693484fc66b15640bc7d17bed16d)) - Tom Oram
- Output ARN to debug - ([86e9c6a](https://github.com/armakuni/terraform-aws-static-website-bucket/commit/86e9c6a03bb8b4d9af9f2b04c08e14e5c3427b60)) - Tom Oram

- - -

## [v0.3.0](https://github.com/armakuni/terraform-aws-static-website-bucket/compare/v0.2.0..v0.3.0) - 2023-09-18
#### Features
- Import module - ([f021ed5](https://github.com/armakuni/terraform-aws-static-website-bucket/commit/f021ed5ccb9f9f128c03c106da08929ea2da9d8a)) - Tom Oram

- - -

## [v0.2.0](https://github.com/armakuni/terraform-aws-static-website-bucket/compare/v0.1.0..v0.2.0) - 2023-09-18
#### Features
- Import module - ([83bc4c4](https://github.com/armakuni/terraform-aws-static-website-bucket/commit/83bc4c4673b3c5c40ab4e3feac8523c28fa6c550)) - Tom Oram

- - -

Changelog generated by [cocogitto](https://github.com/cocogitto/cocogitto).