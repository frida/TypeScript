namespace ts {
    export function prebindSourceFile(file: SourceFile, options: CompilerOptions) {
        bindSourceFile(file, options);
        file.isPrebound = true;
    }
}
