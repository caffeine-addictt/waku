import { SpawnOptionsWithoutStdio, spawn } from 'child_process';

export const execute = <T extends Buffer>(
  command: string,
  args?: string[],
  spawnOptions?: SpawnOptionsWithoutStdio,
): Promise<T> =>
  new Promise((resolve, reject): void => {
    const cmd = spawn(command, args, spawnOptions);

    cmd.stdout.on('data', resolve);
    cmd.stderr.on('data', reject);
  });
